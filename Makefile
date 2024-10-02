## Uso: go-1000 para criar a pasta para o desafio numero 1000
go-%:
	@echo 'Iniciando o desafio $*'
	@cp -r boilerplates/go/ $*
	@echo '| 🛠️     | [$*](https://judge.beecrowd.com/pt/problems/view/$*) | ???      | ???' >> README.md
	@git add README.md


include show-help-minified.make
