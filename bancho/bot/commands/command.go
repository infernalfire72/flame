package commands

import (
	"io/ioutil"
	"strings"
	"sync"

	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func init() {
	// LoadCommands()
}

const Prefix string = "!"

var (
	Commands map[string]*Command
	Mutex    sync.RWMutex
)

type CommandInfo struct {
	Name    string
	Aliases []string
	Syntax  string
}

type Command struct {
	CommandInfo
	Handler func(*objects.Player, []string, objects.Target)
}

func GetCommand(name string) *Command {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return Commands[name]
}

func LoadCommand(path, name string) {
	/*fs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error(err)
		return
	}

	var dat []byte
	for _, file := range fs {
		if strings.HasSuffix(file.Name(), ".go") {
			dat, err = ioutil.ReadFile(path + file.Name())
			if err != nil {
				log.Error(err)
				return
			}

			break
		}
	}

	if len(dat) == 0 {
		log.Error(errors.New("Empty File in Command Folder " + path))
		return
	}

	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	_, err = i.Eval(string(dat))
	if err != nil {
		log.Error(err)
		return
	}

	fn, err := i.Eval("icmd.Run")
	if err != nil {
		log.Error(err)
		return
	}

	syntax, err := i.Eval("icmd.Syntax")
	if err != nil {
		log.Error(err)
		return
	}

	aliases, err := i.Eval("icmd.Aliases")
	if err != nil {
		log.Error(err)
		return
	}

	c := &Command{
		CommandInfo: CommandInfo{
			Name:    name,
			Syntax:  syntax.Interface().(string),
			Aliases: aliases.Interface().([]string),
		},
		Handler: fn.Interface().(func(*objects.Player, []string, objects.Target)),
	}

	Mutex.Lock()
	Commands[name] = c
	Mutex.Unlock()*/
}

func UnloadCommands() {
	Mutex.Lock()
	Commands = map[string]*Command{
		"reload": &Command{
			CommandInfo: CommandInfo{
				Name:    "reload",
				Syntax:  "reload",
				Aliases: []string{"r"},
			},
			Handler: reload,
		},
	}
	Mutex.Unlock()
}

func reload(p *objects.Player, args []string, target objects.Target) {
	LoadCommands()

	target.Write(packets.IrcMessageArgs("system", "Done", target.GetName(), 999))
}

func LoadCommands() {
	UnloadCommands()

	files, err := ioutil.ReadDir("./bancho/bot/commands/")
	if err != nil {
		log.Error(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			fName := f.Name()
			log.Info("Loading Command", fName)

			LoadCommand("./bancho/bot/commands/"+fName+"/", fName)
		}
	}
}

func Execute(sender *objects.Player, content string, target objects.Target) {
	content = content[1:]

	var (
		name string
		args []string
	)

	if split := strings.IndexRune(content, ' '); split != -1 {
		args = strings.Split(content, " ")
		name = args[0]
		args = args[1:]
	} else {
		name = content
	}

	if len(name) == 0 {
		return
	}

	cmd := GetCommand(name)
	if cmd == nil {
		return
	}

	cmd.Handler(sender, args, target)
}
