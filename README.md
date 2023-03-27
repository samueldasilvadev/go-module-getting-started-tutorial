# Go trip

Esse projeto faz parte da série de tutoriais iniciais da Golang presentes
na documentação em: https://go.dev/doc/.

## Criar e usar módulos

```
# inicializa o modulo (go.mod)

go mod init *module-name*

# adiciona pacotes ou remove os que não estão sendo usados

go mod tidy

# muda o apontamento para um pacote local
# opção interessante para testar a implementação de pacotes existentes
# ou que eu estiver desenvolvendo

go mod edit -replace *package-name*=*local-path*
```
