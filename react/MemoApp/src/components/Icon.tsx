import { createIconSetFromIcoMoon } from "@expo/vector-icons"
import { useFonts } from "expo-font"
import fontData from "../../assets/fonts/icomoon.ttf"
import fontSelection from "../../assets/fonts/selection.json"

const CustomIcon = createIconSetFromIcoMoon(
  fontSelection,
  "IconMoon",
  "icomoon.ttf"
)

interface props {
  name: string
  size: number
  color: string
}

export const Icon = (props: props) => {
  const { name, size, color } = props
  const [fontLoaded] = useFonts({
    IconMoon: fontData
  })
  if (!fontLoaded) {
    return null
  }
  return (
    <CustomIcon name={name} size={size} color={color} />
  )
}
