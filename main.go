package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"website/tictactoe"
)

func main() {
	game := tictactoe.NewGame()
	handler := func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			if request.FormValue("reset") == "true" {
				game = tictactoe.NewGame()
				fmt.Printf("Game reset.")
			} else {
				request.ParseForm()
				button := request.FormValue("button")
				fmt.Printf("Button %s clicked\n", button)

				buttonInt, _ := strconv.Atoi(button)
				x, y := (buttonInt-1)/3, (buttonInt-1)%3
				game.MakeMove(x, y)

				winner := game.CheckWinner()
				if winner != "" {
					fmt.Printf("Player %s wins!", winner)
				}
				if game.IsDraw() {
					fmt.Printf("It's a draw!")
				}
			}
		}

		tmpl, err := template.New("index").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Go Web App</title>
				<link rel="stylesheet" type="text/css" href="/static/styles.css">
			</head>
			<body>
				<h1>Tic-Tac-Toe</h1>
				<form method="post">
					<table>
						<tr>
							<td><button name="button" value="1">{{index .Board 0 0}}</button></td>
							<td><button name="button" value="2">{{index .Board 0 1}}</button></td>
							<td><button name="button" value="3">{{index .Board 0 2}}</button></td>
						</tr>
						<tr>
							<td><button name="button" value="4">{{index .Board 1 0}}</button></td>
							<td><button name="button" value="5">{{index .Board 1 1}}</button></td>
							<td><button name="button" value="6">{{index .Board 1 2}}</button></td>
						</tr>
						<tr>
							<td><button name="button" value="7">{{index .Board 2 0}}</button></td>
							<td><button name="button" value="8">{{index .Board 2 1}}</button></td>
							<td><button name="button" value="9">{{index .Board 2 2}}</button></td>
						</tr>
					</table>
					<button type="submit" name="reset" value="true">Reset</button>
				</form>
				{{if .Winner}}
				<h2>{{.Winner}} wins!</h2>
				{{else if .Draw}}
				<h2>It's a draw!</h2>
				{{end}}
			</body>
			</html>
		`)

		if err != nil {
			fmt.Fprintf(writer, "Error: %s", err)
			return
		}

		data := struct {
			Board  [3][3]string
			Winner string
			Draw   bool
		}{
			Board:  game.Board(),
			Winner: game.CheckWinner(),
			Draw:   game.IsDraw(),
		}

		tmpl.Execute(writer, data)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
