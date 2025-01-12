import { FlatList, SafeAreaView, StyleSheet } from 'react-native';
import { ListItem } from '../components/ListItem';
import { useState, useEffect } from "react";
import axios from "axios";
import { useNavigation } from '@react-navigation/native';

const url = "https://newsapi.org/v2/top-headlines?country=us&apiKey=hogehoge"

export const HomeScreen = () => {
  const navigation = useNavigation();

  // [stateの変数, stateを変更する関数] = useState(初期値)
  const [articles, setArticles] = useState([]);

  const fetchArticles = async () => {
    try {
      const response = await axios.get(url)
      setArticles(response.data.articles)
    } catch (error) {
      console.error(error);
    }
  }

  // useEffectの第2引数に空の配列を渡すと初回レンダリング時のみ実行される
  useEffect(() => {
    fetchArticles();
  }, []);

  return (
    <SafeAreaView style={styles.container}>
      <FlatList
        data={articles}
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
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#eee",
  },
});
