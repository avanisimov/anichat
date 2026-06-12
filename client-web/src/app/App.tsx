import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom"
import { useSelector } from "react-redux"

import AuthStart from "../pages/AuthStart"
import Otp from "../pages/Otp"
import ProfileSetup from "../pages/ProfileSetup"
import MainApp from "../pages/MainApp"

import type { RootState } from "../store/store"

export default function App() {
  const isAuth = useSelector((state: RootState) => state.auth.isAuthenticated)
  const isNewUser = useSelector((state: RootState) => state.user.isNewUser)

  return (
    <BrowserRouter>
      <Routes>

        {/* публичные маршруты */}
        {!isAuth && (
          <>
            <Route path="/" element={<AuthStart />} />
            <Route path="/otp" element={<Otp />} />
          </>
        )}

        {/* onboarding нового пользователя */}
        {isAuth && isNewUser && (
          <Route path="/profile" element={<ProfileSetup />} />
        )}

        {/* главное приложение */}
        {isAuth && !isNewUser && (
          <Route path="/app" element={<MainApp />} />
        )}

        {/* fallback */}
        <Route path="*" element={<Navigate to="/" />} />

       

      </Routes>
    </BrowserRouter>
  )
}