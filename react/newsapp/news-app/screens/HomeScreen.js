import { FlatList, SafeAreaView, StyleSheet, RefreshControl } from "react-native";
import { ListItem } from "../components/ListItem";
import { useState, useEffect, useRef } from "react";
import axios from "axios";
import { useNavigation } from "@react-navigation/native";

const url =
  "https://newsapi.org/v2/top-headlines?country=us&apiKey=hogehoge";

export const HomeScreen = () => {
  const navigation = useNavigation();

  // [stateの変数, stateを変更する関数] = useState(初期値)
  const [articles, setArticles] = useState([]);
  const [refreshing, setRefreshing] = useState(false);
  // useRefはコンポーネントの再レンダリングをトリガーしない
  const pageRef = useRef(1);
  const fechedAllRef = useRef(false);

  const fetchArticles = async (page) => {
    try {
      const response = await axios.get(`${url}&page=${page}`);
      if (response.data.articles.length > 0) {
        setArticles([...articles, ...response.data.articles]);
      } else {
        fechedAllRef.current = true;
      }
    } catch (error) {
      console.error(error);
    }
  };

  // useEffectの第2引数に空の配列を渡すと初回レンダリング時のみ実行される
  useEffect(() => {
    fetchArticles(1);
  }, []);

  const onEndReached = () => {
    if (fechedAllRef.current) {
      console.log("onEndReached");
      return;
    }
    pageRef.current = pageRef.current + 1;
    fetchArticles(pageRef.current);
  };

  const onRefresh = async () => {
    setRefreshing(true);
    setArticles([]);
    pageRef.current = 1;
    fechedAllRef.current = false;
    await fetchArticles(1);
    setRefreshing(false);
  };

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
        onEndReached={onEndReached}
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
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
