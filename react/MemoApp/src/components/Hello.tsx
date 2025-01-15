import { View, Text, StyleSheet } from "react-native"

interface props {
  children: string,
  bang?: boolean
}

export const Hello = (props: props) => {
  const { children, bang } = props
  return (
    <View >
      <Text style={styles.text}>Hello, {children} {bang ? "!" : ""}</Text>
    </View >
  )
}

const styles = StyleSheet.create({
  text: {
    color: '#ffffff',
    backgroundColor: "blue",
    fontSize: 40,
    fontWeight: "bold",
    padding: 16
  }
})
