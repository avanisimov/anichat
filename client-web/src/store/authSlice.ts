import { createSlice, PayloadAction } from "@reduxjs/toolkit"

interface AuthState {
  email: string
  token: string | null
  isAuthenticated: boolean
  otpSent: boolean
}

const initialState: AuthState = {
  email: "",
  token: null,
  isAuthenticated: false,
  otpSent: false,
}

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setEmail(state, action: PayloadAction<string>) {
      state.email = action.payload
    },
    otpSentSuccess(state) {
      state.otpSent = true
    },
    loginSuccess(state, action: PayloadAction<string>) {
      state.token = action.payload
      state.isAuthenticated = true
    },
    logout(state) {
      state.token = null
      state.isAuthenticated = false
    },
  },
})

export const { setEmail, otpSentSuccess, loginSuccess, logout } =
  authSlice.actions

export default authSlice.reducer