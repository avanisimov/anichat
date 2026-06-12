import { useRef, useState } from "react"
import { useNavigate } from "react-router-dom"
import { useAppDispatch } from "../store/hooks"
import { loginSuccess } from "../store/authSlice"
import { setNewUser } from "../store/userSlice"

import "./Otp.css"

export default function Otp() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const [code, setCode] = useState([
    "",
    "",
    "",
    "",
    "",
    "",
  ])

  const inputsRef = useRef<(HTMLInputElement | null)[]>([])

  const handleChange = (
    value: string,
    index: number
  ) => {
    if (!/^\d?$/.test(value)) return

    const newCode = [...code]
    newCode[index] = value
    setCode(newCode)

    if (value && index < 5) {
      inputsRef.current[index + 1]?.focus()
    }
  }

  const handleVerify = () => {
    const otp = code.join("")

    if (otp.length !== 6) {
      alert("Введите код полностью")
      return
    }

    console.log("OTP:", otp)

    dispatch(loginSuccess("token_123")) // после создания профиля
     dispatch(setNewUser(true))
    // TODO API verify OTP

    navigate("/profile")
  }

  return (
    <div className="otp-page">
      <div className="otp-card">
        <div className="otp-logo">
          A
        </div>

        <h1 className="otp-title">
          Проверка Email
        </h1>

        <p className="otp-subtitle">
          Мы отправили код подтверждения на
          <br />
          <span className="otp-email">
            example@email.com
          </span>
        </p>

        <div className="otp-inputs">
          {code.map((digit, index) => (
            <input
              key={index}
              ref={(el) => {
                inputsRef.current[index] = el
              }}
              className="otp-input"
              value={digit}
              maxLength={1}
              onChange={(e) =>
                handleChange(
                  e.target.value,
                  index
                )
              }
            />
          ))}
        </div>

        <button
          className="otp-button"
          onClick={handleVerify}
        >
          Подтвердить
        </button>

        <div className="otp-resend">
          <button>
            Отправить код повторно
          </button>
        </div>
      </div>
    </div>
  )
}