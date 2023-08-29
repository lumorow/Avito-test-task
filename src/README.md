# Описание и пояснения
## Инструкция по запуску
*Поднятие и развертывание dev-среды в docker*
>Make run

*Запуск тестов*
>Make tests

*Удаление среды*
>Make сlean

## Список реализованных методов:

*Метод добавления сегмента*
>POST("/segment")

*Метод удаления сегмента*
>DELETE("/segment/:slug")

*Метод добавления пользователя в сегменты*
>PUT("/user/:uid/segments")

*Метод удаления у пользователя сегменты*
>DELETE("/user/:uid/segments")

*Метод получения активных сегментов пользователя*
>GET("/user/:uid/segments")

*Метод удаления пользователя*
>DELETE("/user/:uid")

## Примеры запросов/ответов
-   Пример №1
    - Запрос
    > POST "http://localhost:8080/segment"  
    Body: {  
         segment_name: "AVITO_SUPER_SALE_50",  
    }
    - Ответ

    > {  
        "message": "Segment created successfully",  
        "segment id: 1": {  
            "segment_name": "AVITO_WORK"  
        }  
    }  