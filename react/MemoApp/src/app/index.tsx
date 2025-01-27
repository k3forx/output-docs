import { Redirect, router } from "expo-router"
import { onAuthStateChanged } from "firebase/auth"
import { auth } from "../config"
import { useEffect } from "react"

const Index = () => {
  useEffect(() => {
    onAuthStateChanged(auth, (user) => {
      if (user !== null) {
        router.replace("/memo/list")
      }
    })
  }, [])
  return <Redirect href="/auth/log_in" />
}

export default Index
