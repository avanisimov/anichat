import { createSlice, PayloadAction } from "@reduxjs/toolkit"

interface UserState {
  firstName: string
  lastName: string
  avatar: string | null
  isNewUser: boolean
}

const initialState: UserState = {
  firstName: "",
  lastName: "",
  avatar: null,
  isNewUser: false,
}

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setProfile(state, action: PayloadAction<Partial<UserState>>) {
      Object.assign(state, action.payload)
    },
    setNewUser(state, action: PayloadAction<boolean>) {
      state.isNewUser = action.payload
    },
  },
})

export const { setProfile, setNewUser } = userSlice.actions
export default userSlice.reducer