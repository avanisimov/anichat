import { Navigate } from "react-router-dom"
import { useAppSelector } from "../store/hooks"

interface Props {
  children: React.ReactNode
}

export default function GuestRoute({
  children,
}: Props) {
  const isAuthenticated = useAppSelector(
    (state) => state.auth.isAuthenticated
  )

  const user = useAppSelector(
    (state) => state.user
  )

  const isProfileCompleted =
    !!user.firstName &&
    !!user.lastName

  if (isAuthenticated) {
    return (
      <Navigate
        to={
          isProfileCompleted
            ? "/"
            : "/profile"
        }
        replace
      />
    )
  }

  return <>{children}</>
}