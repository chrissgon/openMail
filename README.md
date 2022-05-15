![OpenMail](https://i.ibb.co/K2pshdw/openmail.png)
# OpenMail
Aplicação de envio de email criada com Vue + Golang + Docker.


## Clonar Repositório
Inicialmente é necessário clonar o repositório principal do projeto com
`git clone git@github.com:chrissgon/openMail.git`.

## Conexão SMTP
Para o envio é necessário criar uma conexão SMTP, por isso informe as credenciais no arquivo `.env` dentro de `api`.
<br>
Por padrão no `.env` o `HOST` e `PORT` preenchidos são do Gmail, por isso use uma conta Gmail, sendo necessário apenas indicar os valores `USERNAME` e `PASSWORD` e habilitar a opção <a target="_blank" href="https://myaccount.google.com/lesssecureapps">Apps Menos Seguros</a>.

## Execução dos Ambientes
Para executar a aplicação use o comando `docker-compose up -d --build`.
<br>
Em seguida acesse <a target="_blank" href="http://localhost:8080/">http://localhost:8080/</a>.