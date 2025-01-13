import { SafeAreaView, StyleSheet, FlatList } from "react-native";
import { useSelector } from "react-redux";
import { ListItem } from "../components/ListItem";
import { useNavigation } from "@react-navigation/native";

export const ClipScreen = () => {
  const clips = useSelector((state) => state.user.clips);
  const navigation = useNavigation();
  return (
    <SafeAreaView style={styles.container}>
      <FlatList
        data={clips}
        renderItem={({ item }) => (
          <ListItem
            imageUrl={item.urlToImage}
            title={item.title}
            author={item.author}
            onPress={() => navigation.navigate("Article", { article: item })}
          />
        )}
        keyExtractor={(_, index) => index.toString()}
      />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#eee",
  },
});
