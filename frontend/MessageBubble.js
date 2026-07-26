import React from 'react';
import {View, Text, StyleSheet} from 'react-native';

export default function MessageBubble({item}) {
  const {type, payload} = item;
  return (
    <View style={styles.bubble}>
      <Text>{type}</Text>
      <Text>{payload}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  bubble: {padding: 8, marginVertical: 2},
});
