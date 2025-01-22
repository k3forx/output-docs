import React from "react"
import { View, StyleSheet } from "react-native"
import { Header } from "../../components/Header"
import { MemoListItem } from "../../components/MemoListItem"
import { CircleButton } from "../../components/CircleButton"
import { Icon } from "../../components/Icon"
import { router } from "expo-router"

const List = () => {
  const handlePress = () => {
    router.push("/memo/create")
  }
  return (
    <View style={styles.container}>
      <Header />
      <View >
        <MemoListItem />
        <MemoListItem />
        <MemoListItem />
      </View>
      <CircleButton onPress={handlePress}>
        <Icon name="plus" size={40} color="#ffffff" />
      </CircleButton>
    </View >
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#ffffff"
  }
})

export default List
