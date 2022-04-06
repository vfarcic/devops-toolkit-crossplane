package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive walkthrough",
	Long:  `Interactive walkthrough`,
	Run: func(cmd *cobra.Command, args []string) {
		m := initialModel()
		if err := tea.NewProgram(m).Start(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

const (
	listHeight        = 14
	defaultWidth      = 20
	titleTypes        = "Which type would you like to use?"
	titleCompositions = "Which Composition would you like to use?"
)

type item string

type model struct {
	xrdApis             list.Model
	quitting            bool
	selectedXrdText     string
	selectedXrdKind     string
	selectedXrdApi      string
	selectedXrdVersion  string
	selectedXrdName     string
	selectedComposition string
	selectedType        string
	types               list.Model
	compositions        list.Model
	screen              string
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

func initialModel() model {
	return model{
		xrdApis:      getXrdApiList(),
		types:        getTypeList(),
		compositions: getCompositionList("", ""),
		screen:       "xrds",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.xrdApis.SetWidth(msg.Width)
		m.types.SetWidth(msg.Width)
		m.compositions.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			switch m.screen {
			case "xrds":
				i, ok := m.xrdApis.SelectedItem().(item)
				if !ok {
					return m, nil
				}
				m.selectedXrdText = string(i)
				m.selectedXrdKind = strings.Split(m.selectedXrdText, " ")[0]
				fullApi := strings.Split(strings.Split(m.selectedXrdText, "(")[1], ")")[0]
				m.selectedXrdName = strings.Split(fullApi, ".")[0]
				apiWithoutName := strings.ReplaceAll(fullApi, m.selectedXrdName+".", "")
				m.selectedXrdApi = strings.Split(apiWithoutName, "/")[0]
				m.selectedXrdVersion = strings.Split(apiWithoutName, "/")[1]
				m.compositions = getCompositionList(m.selectedXrdKind, m.selectedXrdApi+"/"+m.selectedXrdVersion)
				m.screen = "compositions"
				m.compositions.Title = "RXD: " + m.selectedXrdText + "\n\n" + titleCompositions
			case "compositions":
				i, ok := m.compositions.SelectedItem().(item)
				if !ok {
					return m, nil
				}
				m.selectedComposition = string(i)
				m.screen = "types"
				m.types.Title = "RXD: " + m.selectedXrdText + "\n\n" + titleTypes
			case "types":
				i, ok := m.types.SelectedItem().(item)
				if !ok {
					return m, nil
				}
				m.selectedType = string(i)
				m.screen = "explain"
			default:
				m.screen = "end"
			}
			return m, nil
		}
	}
	switch m.screen {
	case "compositions":
		m.compositions, cmd = m.compositions.Update(msg)
	case "types":
		m.types, cmd = m.types.Update(msg)
	default:
		m.xrdApis, cmd = m.xrdApis.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("See ya!")
	}
	switch m.screen {
	case "compositions":
		return m.compositions.View()
	case "types":
		return m.types.View()
	case "explain":
		crd := getCRD(m.selectedXrdName + "." + m.selectedXrdApi)
		xr := getXR(crd, m.selectedComposition)
		yamlData := getXRYaml(xr)
		out := fmt.Sprintf(`XRD: %s
Composition: %s
Type: %s

Sample YAML:

---

%s

---
`,
			m.selectedXrdText,
			m.selectedComposition,
			m.selectedType,
			yamlData,
		)
		return out
	case "end":
		return `
TODO:

* Fix the option to choose claims
* Edit fields
* Export to YAML
* kubectl apply
* Push to Git
* See all running claims, compositions, and managed resources
* Sleep
		
The End`
	}
	return m.xrdApis.View()
}

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

func getXrdApiList() list.Model {
	items := []list.Item{}
	xrds := getXrds()
	for _, xrd := range xrds.Items {
		items = append(items, item(xrd.Spec.Names.Kind+" ("+xrd.Metadata.Name+"/"+xrd.Spec.Versions[0].Name+")"))
	}
	return getListModel(items, "Which Crossplane Resource Definition (XRD) would you like to use?")
}

func getTypeList() list.Model {
	items := []list.Item{item("Composition"), item("Claim")}
	return getListModel(items, titleTypes)
}

func getCompositionList(expectedKind, expectedApi string) list.Model {
	items := []list.Item{}
	compositions := getAllCompositions()
	for _, composition := range compositions.Items {
		if expectedKind == composition.Spec.CompositeTypeRef.Kind &&
			expectedApi == composition.Spec.CompositeTypeRef.ApiVersion {
			items = append(items, item(composition.Metadata.Name))
		}
	}
	return getListModel(items, titleCompositions)
}

func getListModel(items []list.Item, title string) list.Model {
	list := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	list.Title = title
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	list.Styles.Title = titleStyle
	list.Styles.PaginationStyle = paginationStyle
	list.Styles.HelpStyle = helpStyle
	return list
}
