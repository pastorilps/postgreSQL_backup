# Instruções de Uso para o Backup Automático de Banco de Dados PostgreSQL

Este é um script em Go para realizar backups automáticos de bancos de dados PostgreSQL. Ele foi desenvolvido para funcionar com Go 1.20 e PostgreSQL 16. Abaixo estão as instruções para configurar e usar este script em diferentes sistemas operacionais (Windows, Linux e macOS).

## Requisitos

- Go 1.20 ou superior instalado e configurado.
- PostgreSQL 16 ou superior instalado e configurado.
- Permissões de leitura e escrita no diretório onde os backups serão armazenados.

## Configuração do `pg_dump`

Para que o script possa criar backups usando o `pg_dump`, é necessário garantir que o `pg_dump` esteja configurado corretamente no seu ambiente. Aqui estão as instruções para configurar o `pg_dump` em diferentes sistemas operacionais:

### Windows

1. Adicione o diretório binário do PostgreSQL ao PATH do sistema.
2. Abra o Prompt de Comando como administrador.
3. Teste se o `pg_dump` está acessível digitando `pg_dump --version`.

### Linux

1. O `pg_dump` geralmente é instalado junto com o PostgreSQL em sistemas Linux.
2. Verifique se o `pg_dump` está acessível digitando `pg_dump --version` no terminal.

### macOS

1. O `pg_dump` geralmente é instalado junto com o PostgreSQL em sistemas macOS.
2. Verifique se o `pg_dump` está acessível digitando `pg_dump --version` no Terminal.

## Configuração do Arquivo `.env`

Antes de executar o script, certifique-se de configurar as variáveis de ambiente necessárias no arquivo `.env`. Um exemplo de arquivo `.env` é fornecido para você preencher com suas próprias informações. As variáveis de ambiente incluem:

- `HOST_DATABASE`: Endereço do host do banco de dados PostgreSQL.
- `PORT_DATABASE`: Porta do banco de dados PostgreSQL.
- `USER_DATABASE`: Nome de usuário do banco de dados PostgreSQL.
- `PASS_DATABASE`: Senha do usuário do banco de dados PostgreSQL.
- Nomes dos bancos de dados de produção e de playground.

Copie o arquivo `.env.example` e renomeie-o para `.env`, em seguida, preencha com suas próprias credenciais e informações do banco de dados.

## Executando o Script

Após configurar o `pg_dump` e preencher o arquivo `.env`, você está pronto para executar o script. Basta executar o seguinte comando no terminal:

```bash
go run main.go

Isso iniciará o script, que agendará automaticamente backups para serem executados diariamente às 22h (horário do Brasil). Certifique-se de manter o script em execução para que os backups sejam agendados e executados conforme programado.