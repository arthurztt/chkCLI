# chkCLI - Ferramenta para Checagem de Endpoints HTTP

`chkCLI` é uma ferramenta de linha de comando desenvolvida em Go para realizar verificações rápidas em endpoints HTTP. Seu objetivo é fornecer uma maneira simples de validar a disponibilidade de serviços web, retornando o código de status HTTP e o tempo de resposta de cada requisição.

Ideal para testes rápidos de APIs, monitoramento manual de serviços, validação de ambientes e troubleshooting durante o desenvolvimento.

---

## Requisitos

### Executando pelo código-fonte

- Go 1.24 ou superior
- Git

### Executando pelo binário

Nenhum requisito adicional. Basta baixar o executável já compilado, disponibilizado automaticamente na página de **Releases** do projeto a cada nova versão publicada.

---

## Estrutura do Projeto

```
chkCLI/
├── cmd/
│   └── root.go              # Implementação dos comandos e lógica principal
│
├── dist/                    # Artefatos gerados automaticamente a cada release
│   ├── chk_windows_amd64.zip
│   ├── chk_linux_amd64.tar.gz
│   ├── chk_darwin_arm64.tar.gz
│   ├── checksums
│   └── ...
│
├── main.go                  # Ponto de entrada da aplicação
├── go.mod                   # Dependências do projeto
├── go.sum                   # Controle de versões das dependências
├── .goreleaser.yaml         # Configuração do processo de build e release
└── README.md
```

---

## Resumo do Projeto

O **chkCLI** foi desenvolvido utilizando a linguagem **Go** com foco em simplicidade, desempenho e portabilidade.

A aplicação utiliza a biblioteca padrão `net/http` para realizar requisições HTTP e mede o tempo de resposta de cada endpoint utilizando o pacote `time`. O usuário pode informar um ou vários endereços na linha de comando, permitindo verificar diversos serviços em uma única execução.

Cada requisição retorna:

- Código de status HTTP (200, 404, 500, etc.);
- Tempo total da requisição;
- URL consultada.

Também é possível definir um tempo máximo de espera utilizando a flag `--timeout` (`-t`). Caso o endpoint ultrapasse esse limite, a requisição é encerrada e a aplicação informa o erro correspondente.

A interface de linha de comando foi construída utilizando a biblioteca **Cobra**, facilitando a criação de comandos, argumentos e flags de maneira organizada e extensível.

Cada nova versão do projeto é compilada e publicada automaticamente para Windows, Linux e macOS (amd64 e arm64) assim que uma tag é criada no repositório, garantindo que o binário disponível nas Releases esteja sempre atualizado, com checksums de verificação e changelog gerados junto — sem passos manuais de empacotamento.

## Como Executar

### Executando localmente

Clone o repositório:

```bash
git clone https://github.com/arthurztt/chkCLI.git
```

Entre na pasta do projeto:

```bash
cd chkCLI
```

Execute diretamente com o Go:

```bash
go run . https://google.com
```

ou compile o projeto:

```bash
go build -o chk
```

Após a compilação:

Linux/macOS

```bash
./chk https://google.com
```

Windows

```powershell
chk.exe https://google.com
```

Também é possível verificar múltiplos endpoints simultaneamente:

```bash
chk https://google.com https://github.com https://example.com
```

Definindo um timeout personalizado:

```bash
chk --timeout 10 https://google.com
```

ou

```bash
chk -t 10 https://google.com
```

---

### Executando utilizando o binário

Acesse a página de **Releases** do projeto no GitHub e faça o download do executável correspondente ao seu sistema operacional.

Após extrair o arquivo:

Windows

```powershell
chk.exe https://google.com
```

Linux

```bash
./chk https://google.com
```

macOS

```bash
./chk https://google.com
```

Também é possível informar vários endpoints:

```bash
chk https://site1.com https://site2.com https://site3.com
```

---

### Executando como um comando nativo (recomendado)

Para não precisar digitar o caminho completo do binário toda vez, você pode adicioná-lo às variáveis de ambiente do sistema (`PATH`). Assim, o comando `chk` fica disponível globalmente no terminal, de qualquer diretório.

#### Windows

1. Baixe o binário na página de **Releases** e extraia o arquivo `chk.exe` para uma pasta fixa, por exemplo:
   ```
   C:\dev\chkCLI\
   ```
2. Abra o menu Iniciar e pesquise por **"Variáveis de ambiente"** → clique em **"Editar as variáveis de ambiente do sistema"**.
3. Na aba **Avançado**, clique em **Variáveis de Ambiente**.
4. Em **Variáveis do sistema** (ou **Variáveis do usuário**), selecione a variável `Path` e clique em **Editar**.
5. Clique em **Novo** e adicione o caminho da pasta onde está o `chk.exe`:
   ```
   C:\dev\chkCLI\
   ```
6. Confirme em **OK** em todas as janelas e abra um novo terminal (CMD ou PowerShell).
7. Teste digitando:
   ```powershell
   chk https://google.com
   ```

#### Linux/macOS

1. Baixe e extraia o binário, e mova-o para uma pasta que já esteja no `PATH`, como `/usr/local/bin`:
   ```bash
   sudo mv chk /usr/local/bin/chk
   sudo chmod +x /usr/local/bin/chk
   ```
2. Teste digitando, de qualquer diretório:
   ```bash
   chk https://google.com
   ```

Se preferir manter o binário em outra pasta, adicione o caminho ao `PATH` editando o arquivo de configuração do seu shell (`~/.bashrc`, `~/.zshrc`, etc.):

```bash
export PATH="$PATH:/caminho/para/sua/pasta"
```

E recarregue o shell:

```bash
source ~/.bashrc
```

#### macOS (via Homebrew)

Alternativamente, no macOS o `chkCLI` também pode ser instalado diretamente por um tap do Homebrew, sem necessidade de configurar o `PATH` manualmente:

```bash
brew tap arthurztt/homebrew-tap
brew install chk
```

---

## Exemplo de saída

```text
200        87ms      https://google.com
404        42ms      https://example.com/teste
500        151ms     https://api.exemplo.com
```

Caso ocorra timeout ou falha de conexão:

```text
Get "https://site.com": context deadline exceeded   5s   https://site.com
```

---

## Tecnologias Utilizadas

- Go
- Cobra
- net/http
- GoReleaser

---