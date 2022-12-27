import { useRouter } from "next/router"
import { useEffect } from "react"

export default function Home() {

  const router = useRouter()

  useEffect(()=>{
    // We redirect here because we don't have information about index
      router.push("/users")
  }, [])

  return  null
}
