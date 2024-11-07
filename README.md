# DOJO

## Iniciar um novo desafio

Execute o seguinte comando para criar uma pasta do novo desafio

> [!NOTE]
> Você poderá alterar o `1010` para qualquer um dos desafios da [beecrowd](https://judge.beecrowd.com/pt/search)

```bash
make go-1010
```

## Praticando dojo no VS Code

- Instalar [vscode](https://code.visualstudio.com/download)
- Instalar `DojoTools` do `Cristian Dean`
- Ctrl + p (windows / linux) ou cmd + p (mac) para executar o `Dojo Tools: Enable All`

## Começando o desafio

No Beecrowd as entradas do desafio acabam vindo como stdin e o output esperado
é no stout. Abaixo alguns exemplos de como pode começar o código

Utilizando fmt
```go
package main

import (
    "fmt"
)

func main() {
    var a,b int
    var x int

    fmt.Scanf("%d", &a)
    fmt.Scanf("%d", &b)
    x = a + b

    fmt.Printf("X = %d\n", x)
}
```

Utilizando bufio
```go
package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
)

func main() {
    var a,b int
    var x int

    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    a, _ = strconv.Atoi(scanner.Text())

    scanner.Scan()
    b, _ = strconv.Atoi(scanner.Text())

    x = a + b

    fmt.Printf("X = %d\n", x)
}
```

## Problemas

| Status | Problema | Quando | Aonde |
| ------ | -------- | ------ | ----- |
| ☑️      | [2006](https://judge.beecrowd.com/pt/problems/view/2006) | 25/09/24 | [Golang SP & Temporal.io na SumUp](https://www.meetup.com/golangbr/events/303044658)
| ☑️      | [1168](https://judge.beecrowd.com/pt/problems/view/1168) | 08/10/24 | [Dojo com Go na Microsoft Reactor](https://www.meetup.com/golangbr/events/303659869)
| 🕐     | [2164](https://judge.beecrowd.com/pt/problems/view/2164) | ???      | ???
| 🕐     | [1018](https://judge.beecrowd.com/pt/problems/view/1018) | ???      | ???
