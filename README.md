# spotfix

Uma automação para a página do Nexti Web, para lidar com as inconsistência de ponto de mais de 50 mil colaboradores, com um total de 180 mil requisições. Lidando de forma concorrente, utilizando a linguagem Go para eficiência, a biblioteca Rod: sendo versátil, rápida e focada no Chromium. 

Funcionalidades:

- Navegação personalizada
- Ajustes de inconsistências de forma personalizada
- Ajuste de processamento em lote
- Ajuste de processamento com filtro
- CLI para melhor experiência de uso
- Tipo de inconsistências de marcação como "Não registrado", recebem um atraso de 7 dias, para serem processados, para evitar duplicidade do registro.
- Processamento e recusa em lote de todas as inconsistências de marcação diferentes de "Não registrado"

### Instalação(Windows)

Clone o repositório
```bash
git clone https://github.com/isaacgraper/spotfix
```

Instale as dependências
```bash
go mody tidy
```

Crie o build
```bash
go build -o roboCOP .
```

Instale o chromium(Utilizando o chocolatey no Windows)
```bash
choco install chromium
```

### Modo de uso

Listar todos os comandos:
```bash
./roboCOP exec --help
```

Executar processamento em lote sem filtros:
```bash
./roboCOP  exec --batch=10 --max=20 (tamanho do lote: 10, máximo: 20)
```

Executar o processamento em lote, passando hora e tipo da inconsistência:
```bash
./roboCOP exec --filter --hour="9:00" --category="Fora do perímetro"
```

Executar processamento utlizando a filtragem das inconsistências:
```bash
./roboCOP exec --filter
```

Executar processamento utlizando a filtragem das inconsistências, passando o tipo da inconsistência:
```bash
./roboCOP exec --filter --filter-category="Não registrado"
```

### Afazer:

- Modulo de email



