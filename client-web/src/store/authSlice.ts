import { createSlice, PayloadAction } from "@reduxjs/toolkit"

interface AuthState {
  email: string
  accessToken: string | null

  isAuthenticated: boolean
  loading: boolean
}

const accessToken = localStorage.getItem("accessToken")

const initialState: AuthState = {
  email: "",

  accessToken,

  isAuthenticated: !!accessToken,

  loading: true,
}

const authSlice = createSlice({
  name: "auth",

  initialState,

  reducers: {
    setEmail(
      state,
      action: PayloadAction<string>
    ) {
      state.email = action.payload
    },

    setAuthenticated(
      state,
      action: PayloadAction<string>
    ) {
      state.accessToken = action.payload
      state.isAuthenticated = true
    },

    setLoading(
      state,
      action: PayloadAction<boolean>
    ) {
      state.loading = action.payload
    },

    logout(state) {
      state.email = ""

      state.accessToken = null

      state.isAuthenticated = false

      localStorage.removeItem("accessToken")
    },
  },
})

export const {
  setEmail,
  setAuthenticated,
  setLoading,
  logout,
} = authSlice.actions

export default authSlice.reducer