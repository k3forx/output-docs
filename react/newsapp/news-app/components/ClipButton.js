import { StyleSheet, TouchableOpacity } from "react-native";
import { FontAwesome } from "@expo/vector-icons";

export const ClipButton = ({ onPress, enabled }) => {
  const name = enabled ? "bookmark" : "bookmark-o";
  return (
    <TouchableOpacity>
      <FontAwesome
        name={name}
        size={40}
        color="salmon"
        onPress={onPress}
        style={styles.container}
      />
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  container: {
    padding: 5,
  },
});
