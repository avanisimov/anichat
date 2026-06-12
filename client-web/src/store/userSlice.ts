import {
  createSlice,
  PayloadAction,
} from "@reduxjs/toolkit"

interface UserState {
  id: string | null
  email: string | null

  firstName: string | null
  lastName: string | null

  avatar: string | null
}

const initialState: UserState = {
  id: null,
  email: null,

  firstName: null,
  lastName: null,

  avatar: null,
}

const userSlice = createSlice({
  name: "user",

  initialState,

  reducers: {
    setUser(
      state,
      action: PayloadAction<UserState>
    ) {
      return action.payload
    },

    clearUser() {
      return initialState
    },
  },
})

export const {
  setUser,
  clearUser,
} = userSlice.actions

export default userSlice.reducer