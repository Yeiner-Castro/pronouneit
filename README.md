# Pronouneit

## Descripción
Pronouneit es un proyecto desarrollado en Go que utiliza [Gin](https://github.com/gin-gonic/gin) para el manejo de rutas en una aplicación web, y [GORM](https://gorm.io) junto con el driver de MySQL para la interacción con bases de datos. Este proyecto está diseñado para ofrecer una solución eficiente y escalable para aplicaciones web modernas.

## Requisitos Previos
Antes de instalar y ejecutar este proyecto, asegúrate de tener instalado Go en tu sistema. Puedes descargar e instalar Go desde [el sitio web oficial de Go](https://golang.org/dl/).

## Instalación

Para comenzar a usar Pronouneit, primero clona el repositorio en tu máquina local y navega al directorio del proyecto:

!git clone https://github.com/Yeiner-Castro/pronouneit.git
!cd pronouneit

Luego, instala las dependencias necesarias utilizando los siguientes comandos:

!go get -u github.com/gin-gonic/gin
!go get -u gorm.io/gorm
!go get -u gorm.io/driver/mysql

Estos comandos instalarán Gin para el manejo de rutas HTTP, GORM como ORM para manipular la base de datos, y el driver específico de GORM para MySQL.

## Uso
Para iniciar la aplicación, asegúrate de que estás en el directorio principal del proyecto y ejecuta:

