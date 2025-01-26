import {
  View,
  Text,
  TextInput,
  StyleSheet,
  TouchableOpacity
} from "react-native"
import { Button } from "../../components/Button"
import { Link, router } from "expo-router"
import { useState } from "react"

const LogIn = () => {
  const handlePress = () => {
    // ログイン処理
    router.replace("/memo/list")
  }

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  return (
    <View style={styles.container}>
      <View style={styles.inner}>
        <Text style={styles.title}>Log In</Text>
        <TextInput
          style={styles.input}
          value={email}
          onChangeText={setEmail}
          autoCapitalize="none"
          keyboardType="email-address"
          placeholder="メールアドレス"
        />
        <TextInput
          style={styles.input}
          value={password}
          onChangeText={setPassword}
          autoCapitalize="none"
          secureTextEntry
          placeholder="パスワード"
        />

        <Button label="Submit" onPress={handlePress} />

        <View style={styles.footer}>
          <Text style={styles.footerText}>Not registered?</Text>
          <Link href="/auth/sign_up" asChild>
            <TouchableOpacity>
              <Text style={styles.footerLink}>Sign Up here!</Text>
            </TouchableOpacity>
          </Link>
        </View>
      </View>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#F0F4F8"
  },
  inner: {
    paddingVertical: 24,
    paddingHorizontal: 27
  },
  title: {
    fontSize: 24,
    lineHeight: 32,
    fontWeight: "bold",
    marginBottom: 24
  },
  input: {
    borderWidth: 1,
    borderColor: "#DDDDDD",
    backgroundColor: "#ffffff",
    height: 48,
    padding: 8,
    fontSize: 16,
    marginBottom: 16
  },
  footer: {
    flexDirection: "row"
  },
  footerText: {
    fontSize: 14,
    lineHeight: 24,
    marginRight: 8,
    color: "#000000"
  },
  footerLink: {
    fontSize: 14,
    lineHeight: 24,
    color: "#467FD3"
  }
})

export default LogIn
