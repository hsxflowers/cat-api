
# Gatinhos API

Esse repositório tem o objetivo de fonercer uma API com a função de adicionar e puxar diversas imagens de gatinhos de forma randomica.

## Documentação da API

#### Retorna todos os itens

```http
  GET /cat
```

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `---` | `---` | **---** |

#### Retorna um item randomico

```http
  GET /cat/${emotion}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `emotion`      | `string` | **Obrigatório**. A emoção que você deseja que seja retornada |

#### Casos possíveis

- `muitotriste`
- `triste`
- `muitofeliz`
- `feliz`

## Rodando o código

Para rodar o código, após ter Golang instalado, rode o seguinte comando:

```bash
  go run .\cmd\api\main.go
```

