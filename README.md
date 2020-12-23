# go-graphql
GraphQL Tests in Golang

https://github.com/graphql-go/graphql

Permitir query em graphql nos modelos descritos abaixo:
- Um seller tem como atributos possíveis:
    - id: Int64
    - nome: String
- Um seller faz N vendas
- Uma venda tem como atributos possíveis:
    - id: Int64
- Uma venda é composta por N itens
- Um Item tem como atributos possíveis:
    - id: Int64
    - descricao: String
    - valor: Float
    