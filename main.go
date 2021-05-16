package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "os/exec"

var reader *bufio.Reader
type User struct{
	id int
	username string
	email string
	age int
}

var id int
var users map [int]User

func crearUsuario(){
	clearConsole()
	fmt.Print("Ingresa username  ")
	username := readLine()

	fmt.Print("Ingresa email  ")
	email := readLine()

	fmt.Print("Ingresa edad  ")
	age  , err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible converitr a entero")
	}

	id++
	user := User{ id , username ,email ,age}
	users[id] = user
	fmt.Println(users)
	fmt.Println("<<<<< Usuario creado")
}
func listarUsuario(){
	clearConsole()
	for id,user := range users{
		fmt.Println(id, "-" , user.username)
	}

	fmt.Println("Usuario listado")

}
func actualizarUsuario(){
	fmt.Println("Usuario actualizado")

}
func eliminarUsuario(){
	clearConsole()
	fmt.Print("Ingresa el id a eliminar  ")
	id , err := strconv.Atoi(readLine())
	if err != nil {
		panic("No es posible converitr a entero")
	}

	if _, ok:=users[id];ok{
		delete(users, id)
	} 

	fmt.Println(">>>>>Usuario eliminado")

}

func clearConsole (){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {

	if option, err := reader.ReadString('\n') ; err != nil{
		panic("No se puede obytener el valor")
	}else{
		return strings.TrimSuffix(option , "\n")
	}

}

func main(){
	
	var option string 
	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)

	for{
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")
	
		fmt.Print("Ingresa una opcion:  ")
		option = readLine()
		
		if option == "quit" || option == "q"{
			break
		}
		switch option {
		case "a" , "crear":
			crearUsuario()
		case "b" , "listar":
			listarUsuario()
		case "c" , "actualizar":
			actualizarUsuario()
		case "d" , "eliminar":
			eliminarUsuario()
		default :
			fmt.Println("Opcion no valida")	
		}

	}
	fmt.Println("Programa terminado")



}