import { Navigate } from "react-router-dom"
import { useAppSelector } from "../store/hooks"

interface Props {
  children: React.ReactNode
}

export default function ProfileCompletedRoute({
  children,
}: Props) {
  const user = useAppSelector(
    (state) => state.user
  )

  const isProfileCompleted =
    !!user.firstName &&
    !!user.lastName

  if (!isProfileCompleted) {
    return <Navigate to="/profile" replace />
  }

  return <>{children}</>
}