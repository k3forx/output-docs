import { StatusBar } from 'expo-status-bar';
import { StyleSheet, View } from 'react-native';
import { ListItem } from './components/ListItem';

export default function App() {
  const sampleText = "Hogehogeafkljlfjalsjdflaskjdflajsfl asdflkjasldfjal asdlfj lasjf lal jasfljwoiej onbaoeghoavn:ao ij aovnaoig awno aaoi nowij oiaoij:o j jaoj a:oeija :oijaw:ie fjaw:ojf:oawifj:sovjaojf"
  return (
    <View style={styles.container}>
      <ListItem
        imageUrl={"https://picsum.photos/id/10/300/300"}
        title="hogehoge"
        author="React News"
      />
      <ListItem
        imageUrl={"https://picsum.photos/id/20/300/300"}
        title="fugafuga"
        author="Japan News"
      />
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#eee",
    alignItems: "center",
    justifyContent: "center",
  },
});
