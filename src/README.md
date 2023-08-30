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
-   Пример №1. Добавление сегмента
    - Запрос
*       POST "http://localhost:8080/segment"  
        Body: {  
            segment_name: "AVITO_SUPER_SALE_50",  
        }

    - Ответ 
*       "segment id: 1": {  
            "segment_name": "AVITO_SUPER_SALE_50"  
        }
- Пример №2. Добавление сегмента
    - Запрос
*       POST "http://localhost:8080/segment"  
        Body: {  
            segment_name: "AVITO_WOW",  
        }

    - Ответ
*       "segment id: 2": {  
            "segment_name": "AVITO_WOW"  
        }
- Пример №3. Добавление пользователя в сегменты
    - Запрос
*       PUT "http://localhost:8080/user/2000/segments"
        Body: {
            "segments_name": ["AVITO_WOW","AVITO_SUPER_SALE_50"]
        }

    - Ответ
*       {
            "user_id": 2000,
            "segments": [
                 "AVITO_WOW",
                 "AVITO_SUPER_SALE_50"
            ]
        }
- Пример №4. Получение сегментов пользователя
    - Запрос
*       GET "http://localhost:8080/user/2001/segments"

    - Ответ
*       {
            "user_id": 2000,
            "segments": [
                 "AVITO_WOW",
                 "AVITO_SUPER_SALE_50"
            ]
        }
- Пример №5. Удаление сегмента
    - Запрос
*       DELETE "http://localhost:8080/segment/AVITO_SUPER_SALE_50"

    - Ответ
*       {
            "message": "Segment deleted successfully"
        }
- Пример №6. Удаление сегментов у пользователя
    - Запрос
*       DELETE "http://localhost:8080/user/2000/segments"
        {
            "segments_name": ["AVITO_WOW"]
        }
    - Ответ
*       {
            "message": "Success to delete segments to user"
        }
- Пример №7. Удаление пользователя
    - Запрос
*       DELETE "http://localhost:8080/user/2000"

    - Ответ
*       {
            "message": "User deleted successfully"
        }