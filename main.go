package main

import (
    "fmt"
    "os"
    "time"
    "log"
    "os/exec"
    "runtime"
    "strings"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/key"
    tea "github.com/charmbracelet/bubbletea"
)

type Stopwatch struct {
    startTime   time.Time
    elapsedTime time.Duration
    running    bool
}

func (s *Stopwatch) Start() {
    if !s.running {
        s.startTime = time.Now()
        s.running = true
    }
}

func (s *Stopwatch) Stop() {
    if s.running {
        s.elapsedTime += time.Since(s.startTime)
        s.running = false
    }
}

func (s *Stopwatch) Reset() {
    s.startTime = time.Time{}
    s.elapsedTime = 0
    s.running = false
}

func (s *Stopwatch) Format() string {
    elapsed := s.Elapsed()
    hours := int(elapsed.Hours())
    minutes := int(elapsed.Minutes()) % 60
    seconds := int(elapsed.Seconds()) % 60
    return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func (s *Stopwatch) Elapsed() time.Duration {
    if s.running {
        return s.elapsedTime + time.Since(s.startTime)
    }
    return s.elapsedTime
}

type model struct {
    status       string
    stopwatch    Stopwatch
    tracking     bool
    confirmSave  bool
    keys         keyMap
    width        int
    height       int
}

type keyMap struct {
    start   key.Binding
    stop    key.Binding
    confirm key.Binding
    cancel  key.Binding
    quit    key.Binding
}

func defaultKeyMap() keyMap {
    return keyMap{
        start: key.NewBinding(
            key.WithKeys("s"),
            key.WithHelp("s", "start timer"),
        ),
        stop: key.NewBinding(
            key.WithKeys("x"),
            key.WithHelp("x", "stop timer"),
        ),
        confirm: key.NewBinding(
            key.WithKeys("y"),
            key.WithHelp("y", "confirm"),
        ),
        cancel: key.NewBinding(
            key.WithKeys("n"),
            key.WithHelp("n", "cancel"),
        ),
        quit: key.NewBinding(
            key.WithKeys("q", "ctrl+c"),
            key.WithHelp("q", "quit"),
        ),
    }
}

func initialModel() model {
    return model{
        status:     "Welcome to T2!",
        stopwatch:  Stopwatch{},
        tracking:   false,
        keys:       defaultKeyMap(),
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    case tea.KeyMsg:
        switch {
        case key.Matches(msg, m.keys.quit):
            return m, tea.Quit
        case key.Matches(msg, m.keys.start):
            if !m.tracking {
                m.stopwatch.Start()
                m.tracking = true
                m.status = "Timer started. Press 'x' to stop."
            }
        case key.Matches(msg, m.keys.stop):
            if m.tracking {
                m.stopwatch.Stop()
                m.tracking = false
                m.status = fmt.Sprintf("Timer stopped. Total time: %s. Save? (y/n)", m.stopwatch.Format())
                m.confirmSave = true
            }
        case key.Matches(msg, m.keys.confirm):
            if m.confirmSave {
                saveTimeToFile(int(m.stopwatch.Elapsed().Minutes()))
                m.status = "Time saved successfully!"
                return m, tea.Quit
            }
        case key.Matches(msg, m.keys.cancel):
            if m.confirmSave {
                m.status = "Time tracking cancelled"
                return m, tea.Quit
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    softPink := lipgloss.Color("#FFB6C1")
    deepPink := lipgloss.Color("#FF1493")
    
    containerStyle := lipgloss.NewStyle().
        Width(m.width).
        Height(m.height).
        Align(lipgloss.Center, lipgloss.Center).
        Foreground(softPink)

    titleStyle := lipgloss.NewStyle().
        Bold(true).
        Padding(1, 1).
        Align(lipgloss.Center).
        Foreground(deepPink)

    content := []string{
        titleStyle.Render("T2"),
        titleStyle.Copy().Foreground(softPink).Render(m.status),
        "",
        "Press 's' to start timer",
        "Press 'x' to stop timer", 
        "Press 'q' to quit",
    }
    return containerStyle.Render(
        strings.Join(content, "\n"),
    )
}


func clearTerminal() {
    var cmd *exec.Cmd
    switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("cmd", "/c", "cls")
    default:
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func saveTimeToFile(minutes int) {
    file, err := os.OpenFile("tracked-time.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()
    _, err = file.WriteString(fmt.Sprintf("Time tracked: %d minute(s)\n", minutes))
    if err != nil {
        log.Println("Error writing to file:", err)
    }
}

func main() {
    clearTerminal()
    p := tea.NewProgram(
        initialModel(), 
        tea.WithAltScreen(),
        tea.WithMouseCellMotion(),
    )
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
}

