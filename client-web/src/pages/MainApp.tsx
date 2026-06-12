import "./MainApp.css"

export default function MainApp() {
  const chats = [
    {
      id: 1,
      name: "Иван Иванов",
      lastMessage: "Привет 👋",
    },
    {
      id: 2,
      name: "Анна",
      lastMessage: "Как дела?",
    },
    {
      id: 3,
      name: "Петр",
      lastMessage: "Созвонимся вечером",
    },
  ]

  return (
    <div className="main-layout">
      <aside className="sidebar">
        <div className="sidebar-header">
          <div className="logo">
            Anichat
          </div>

          <input
            className="search-input"
            placeholder="Поиск..."
          />
        </div>

        <div className="chat-list">
          {chats.map((chat) => (
            <div
              key={chat.id}
              className="chat-item"
            >
              <div className="chat-avatar">
                {chat.name[0]}
              </div>

              <div className="chat-info">
                <div className="chat-name">
                  {chat.name}
                </div>

                <div className="chat-preview">
                  {chat.lastMessage}
                </div>
              </div>
            </div>
          ))}
        </div>

        <div className="sidebar-footer">
          <button className="settings-button">
            ⚙ Настройки
          </button>
        </div>
      </aside>

      <main className="chat-area">
        <div className="chat-header">
          <div className="chat-header-name">
            Иван Иванов
          </div>
        </div>

        <div className="messages">
          <div className="message">
            Привет!
          </div>

          <div className="message self">
            Привет 👋
          </div>

          <div className="message">
            Как дела?
          </div>
        </div>

        <div className="message-input-area">
          <input
            className="message-input"
            placeholder="Введите сообщение..."
          />

          <button className="send-button">
            ➤
          </button>
        </div>
      </main>
    </div>
  )
}