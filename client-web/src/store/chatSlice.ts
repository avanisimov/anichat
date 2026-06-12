import { createSlice, PayloadAction } from "@reduxjs/toolkit"

interface Message {
  id: string
  text: string
  senderId: string
}

interface ChatState {
  chats: any[]
  activeChatId: string | null
  messages: Record<string, Message[]>
}

const initialState: ChatState = {
  chats: [],
  activeChatId: null,
  messages: {},
}

const chatSlice = createSlice({
  name: "chat",
  initialState,
  reducers: {
    setChats(state, action: PayloadAction<any[]>) {
      state.chats = action.payload
    },
    setActiveChat(state, action: PayloadAction<string>) {
      state.activeChatId = action.payload
    },
    addMessage(
      state,
      action: PayloadAction<{ chatId: string; message: Message }>
    ) {
      const { chatId, message } = action.payload
      if (!state.messages[chatId]) {
        state.messages[chatId] = []
      }
      state.messages[chatId].push(message)
    },
  },
})

export const { setChats, setActiveChat, addMessage } = chatSlice.actions
export default chatSlice.reducer