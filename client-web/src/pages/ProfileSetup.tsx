import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { useAppDispatch } from "../store/hooks"
import { setNewUser } from "../store/userSlice"
import "./ProfileSetup.css"

export default function ProfileSetup() {
const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const [firstName, setFirstName] = useState("")
  const [lastName, setLastName] = useState("")
  const [avatar, setAvatar] = useState<string | null>(null)

  const handleAvatarChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0]

    if (!file) return

    const imageUrl = URL.createObjectURL(file)

    setAvatar(imageUrl)
  }

  const handleSave = () => {
    if (!firstName.trim()) {
      alert("Введите имя")
      return
    }

    console.log({
      firstName,
      lastName,
      avatar,
    })

    // TODO API create profile
    dispatch(setNewUser(false))
    navigate("/app")
  }

  return (
    <div className="profile-page">
      <div className="profile-card">
        <h1 className="profile-title">
          Создание профиля
        </h1>

        <p className="profile-subtitle">
          Расскажите немного о себе
        </p>

        <div className="avatar-wrapper">
          <label className="avatar-upload">
            <div className="avatar-preview">
              {avatar ? (
                <img
                  src={avatar}
                  alt="Avatar"
                />
              ) : (
                <div className="avatar-placeholder">
                  👤
                </div>
              )}
            </div>

            <div className="avatar-overlay">
              +
            </div>

            <input
              className="avatar-input"
              type="file"
              accept="image/*"
              onChange={handleAvatarChange}
            />
          </label>
        </div>

        <label className="profile-label">
          Имя
        </label>

        <input
          className="profile-input"
          placeholder="Введите имя"
          value={firstName}
          onChange={(e) =>
            setFirstName(e.target.value)
          }
        />

        <label className="profile-label">
          Фамилия
        </label>

        <input
          className="profile-input"
          placeholder="Введите фамилию"
          value={lastName}
          onChange={(e) =>
            setLastName(e.target.value)
          }
        />

        <button
          className="profile-button"
          onClick={handleSave}
        >
          Сохранить и продолжить
        </button>
      </div>
    </div>
  )
}