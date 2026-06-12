import { AppDispatch } from "./store"

import {
  setAuthenticated,
  logout,
  setLoading,
} from "./authSlice"

import {
  setUser,
  clearUser,
} from "./userSlice"

export const checkAuth =
  () => async (dispatch: AppDispatch) => {

    const token =
      localStorage.getItem("accessToken")

    if (!token) {
      dispatch(setLoading(false))
      return
    }

    try {
      dispatch(setLoading(true))

      const response = await fetch(
        "http://localhost:8080/api/users/me",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      )

      if (!response.ok) {
        throw new Error("Unauthorized")
      }

      const user = await response.json()

      dispatch(
        setAuthenticated(token)
      )

      dispatch(
        setUser(user)
      )
    } catch (error) {
      localStorage.removeItem(
        "accessToken"
      )

      dispatch(logout())

      dispatch(clearUser())
    } finally {
      dispatch(setLoading(false))
    }
  }