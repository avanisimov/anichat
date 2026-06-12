import {
  Routes,
  Route,
  BrowserRouter,
} from "react-router-dom"

import AuthStart from "../pages/AuthStart"
import Otp from "../pages/Otp"
import ProfileSetup from "../pages/ProfileSetup"
import MainApp from "../pages/MainApp"

import GuestRoute from "./GuestRoute"
import ProtectedRoute from "./ProtectedRoute"

export default function AppRouter() {
  return (
   <BrowserRouter>
        <Routes>

        <Route
            path="/auth"
            element={
            <GuestRoute>
                <AuthStart />
            </GuestRoute>
            }
        />

        <Route
            path="/auth/otp"
            element={
            <GuestRoute>
                <Otp />
            </GuestRoute>
            }
        />

        <Route
            path="/profile"
            element={
            <ProtectedRoute>
                <ProfileSetup />
            </ProtectedRoute>
            }
        />

        <Route
            path="/"
            element={
            <ProtectedRoute>
                <MainApp />
            </ProtectedRoute>
            }
        />

        </Routes>
   </BrowserRouter>
  )
}