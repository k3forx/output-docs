import { View, Text, ScrollView, StyleSheet } from "react-native"
import { CircleButton } from "../../components/CircleButton"
import { Icon } from "../../components/Icon"
import { router } from "expo-router"

const Detail = () => {
  const handlePress = () => {
    router.push("/memo/edit")
  }
  return (
    <View style={styles.container}>
      <View style={styles.memoHeader}>
        <Text style={styles.memoTitle}>Detail</Text>
        <Text style={styles.memoDate}>日付</Text>
      </View>
      <ScrollView style={styles.memoBody}>
        <Text style={styles.memoBodyText}>
          本文 aaaaaaaaaa bbbbbbbbbbbbbbbb
          cccccccccccccllskfjksljaflskjfoawjvoaijefo:ajlsjflaksjdlkajsldfjasldfjalskjdflasjdfoiawjgonabouiearhgaong
        </Text>
      </ScrollView>
      <CircleButton onPress={handlePress} style={{ top: 60, bottom: "auto" }}>
        <Icon name="pencil" size={40} color="white" />
      </CircleButton>
    </View>
  )
}

export default Detail

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#ffffff"
  },
  memoHeader: {
    backgroundColor: "#467FD3",
    height: 96,
    justifyContent: "center",
    paddingVertical: 24,
    paddingHorizontal: 19
  },
  memoTitle: {
    color: "#ffffff",
    fontSize: 20,
    lineHeight: 32,
    fontWeight: "bold"
  },
  memoDate: {
    color: "#ffffff",
    fontSize: 12,
    lineHeight: 16
  },
  memoBody: {
    paddingVertical: 32,
    paddingHorizontal: 27
  },
  memoBodyText: {
    fontSize: 16,
    lineHeight: 24,
    color: "#000000"
  }
})
