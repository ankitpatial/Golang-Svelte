## Channel
A notification channel  

```mermaid
erDiagram
    CHANNEL ||--|{ PARTICEPENTS: have
    CHANNEL ||--o{ MESSAGES: contains
    PARTICEPENTS }|..o{ MESSAGES: have
```
