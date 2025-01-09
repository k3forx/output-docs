import { StatusBar } from 'expo-status-bar';
import { SafeAreaView, ScrollView, StyleSheet } from 'react-native';
import { ListItem } from './components/ListItem';
import articles from './dummies/articles.json';

export default function App() {
  const items = articles.map((article, index) => {
    return (
      <ListItem
        key={index.toString()}
        imageUrl={article.urlToImage}
        title={article.title}
        author={article.author}
      />
    )
  })
  return (
    <SafeAreaView style={styles.container}>
      <ScrollView>
        {items}
      </ScrollView>
      <StatusBar style="auto" />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#eee",
  },
});
