О проекте
Проект представляет собой backend-часть маркетплейса, написанный на языке Go с использованием современных подходов к архитектуре и разработке. Основная цель — создать масштабируемое, безопасное и гибкое API-приложение.
Архитектура
Проект построен по принципам слоистой архитектуры (Clean Architecture):

Handler — слой приёма HTTP-запросов, отвечает за маршруты и работу с входными данными
Service — бизнес-логика приложения, проверка прав, обработка операций
Repository — работа с базой данных через pgx

Такой подход обеспечивает:

Тестируемость: каждый слой может быть протестирован независимо
Поддерживаемость: чёткие границы между компонентами
Масштабируемость: простота расширения и замены отдельных слоёв

Стек технологий

Go (Gin) — высокопроизводительный фреймворк для построения REST API
PostgreSQL + pgx — работа с базой данных на низком уровне, без использования ORM
Redis — кэширование данных, управление сессиями и реализация очередей задач
golang-migrate — миграции и управление схемой базы данных
Viper — гибкая конфигурация окружений
Zap — структурированное логирование с высокой производительностью
JWT — авторизация и разграничение доступа

Текущий статус
Проект находится в стадии активной разработки. Основной функционал постепенно добавляется.
