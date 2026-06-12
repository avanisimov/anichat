import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { useAppDispatch } from "../store/hooks"
import { setEmail } from "../store/authSlice"
import "./AuthStart.css"

export default function AuthStart() {
  const [email, setEmailLocal] = useState("")
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const handleContinue = () => {
    if (!email.includes("@")) {
      alert("Введите корректный email")
      return
    }

    dispatch(setEmail(email))

    // TODO: API запрос отправки OTP

    navigate("/auth/otp")
  }

  return (
    <div className="auth-page">
      <div className="auth-card">
        <div className="auth-logo">A</div>

        <h1 className="auth-title">
          Добро пожаловать в Anichat
        </h1>

        <p className="auth-subtitle">
          Войдите через email для продолжения
        </p>

        <label className="auth-label">
          Email
        </label>

        <input
          className="auth-input"
          type="email"
          placeholder="example@email.com"
          value={email}
          onChange={(e) => setEmailLocal(e.target.value)}
        />

        <button
          className="auth-button"
          onClick={handleContinue}
        >
          Продолжить
        </button>
      </div>
    </div>
  )
}

// export default function AuthStart() {
//   const [email, setEmailLocal] = useState("")
//   const dispatch = useAppDispatch()
//   const navigate = useNavigate()

//   const handleContinue = async () => {
//     if (!email.includes("@")) return alert("Введите корректный email")

//     dispatch(setEmail(email))

//     // TODO: API запрос отправки OTP
//     console.log("Send OTP to:", email)

//     navigate("/otp")
//   }

//   return (
//     <div style={{ padding: 20 }}>
//       <h1>Добро пожаловать в Anichat</h1>

//       <input
//         placeholder="Email"
//         value={email}
//         onChange={(e) => setEmailLocal(e.target.value)}
//       />

//       <button onClick={handleContinue}>
//         Продолжить
//       </button>
//     </div>
//   )
// }
