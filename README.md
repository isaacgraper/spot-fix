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

### Instalação

Clone o repositório
```bash
git clone https://github.com/isaacgraper/spotfix
```

Criar a imagem docker
```bash
docker build -t spotfix-image .
```

Executar a imagem passando a filtragem das inconsistências
```bash
docker run -it --rm spotfix-image --filter
```




