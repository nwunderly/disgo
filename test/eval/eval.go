package eval

import (
	"fmt"
	"github.com/nwunderly/disgo"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"log"
	"os"
	"reflect"
	"strings"
)

var interpreter *interp.Interpreter

//var Symbols = map[string]map[string]reflect.Value{}

func Setup() {
	interpreter = interp.New(interp.Options{GoPath: os.Getenv("GOPATH")})

	stdlib.Symbols["github.com/nwunderly/disgo"] = map[string]reflect.Value{
		"Context": reflect.ValueOf((*disgo.Context)(nil)),
	}

	interpreter.Use(stdlib.Symbols)
	src := `package foo
import (
	"fmt"
	"math"
//	"github.com/bwmarrin/discordgo"
//	"github.com/nwunderly/disgo"
)`
	log.Println("compiling initial eval stuff")
	_, err := interpreter.Eval(src)
	if err != nil {
		fmt.Println(err)
	}
}

func Eval(ctx *disgo.Context) error {
	input := strings.TrimSpace(strings.TrimPrefix(ctx.Message.Content, ctx.Bot.Prefix+"eval"))
	src := fmt.Sprintf(`package foo
func Bar() interface{} {return %s}`, input)

	log.Println("compiling")
	_, err := interpreter.Eval(src)
	if err != nil {
		return err
	}

	log.Println("getting function")
	v, err := interpreter.Eval("foo.Bar")
	if err != nil {
		return err
	}

	f := v.Interface().(func() interface{})

	log.Println("running")
	ctx.Send(fmt.Sprint(f()))
	return nil
}
