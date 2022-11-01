## Библиотеки для прототипа системы b2b2b.

ответвился, чтобы добавить в протокол передачу токена для авторизации

# object
В пакете object описаны основные сущности:
- Entity - организация (юридическое лицо, предприятие)
- Obligation - денежное обязательство
- Origin - происхождение денежного обязательства

# torepo
Пакет, определяющий протокол взаимодействия по gRPC с репозиторием.

# repo_client
Пакет, реализующий функции клиента gRPC для доступа к репозиторию

# repo_srv
Пакет, реализующий Controller, принимающий запрос по gRPC на стороне сервера