# Bastille-Web-V2 (This project is under development...)

> A web interface for the FreeBSD Bastille Jails management. 

## 💻 Requisites

Before begin, you must to know that this project was tested with:

- FreeBSD 14+;
- Golang 1.24+;
- Bastille 1.0.1.250714+;
- ttyd 1.7.7+;

## 🚀 Install bastille-web-v2

To install the bastille-web-v2, follow these steps.

Install dependencies:

```
pkg install -y go bastille ttyd
```

Clone the project and build it:

```
git clone https://github.com/blss-tico/bastille-web-v2.git
cd bastille-web-v2
go build
```

## ☕ Usando bastille-web-v2

To use bastille-web-v2, with default ip/port option:
```
sudo ./bastille-web-v2
```

Open browser and point to machine ip to use the web interface. 

Login user admin and default password admin. Please change default password (admin).

Ip address and port are defined in .env file.

## 📝 Licença

Esse projeto está sob licença. Veja o arquivo [LICENÇA](LICENSE.md) para mais detalhes.
