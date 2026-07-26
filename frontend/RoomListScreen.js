import React from 'react';
import {View, Text, FlatList, TouchableOpacity, StyleSheet} from 'react-native';

export default function RoomListScreen({rooms, onCreate}) {
  return (
    <FlatList
      data={rooms}
      keyExtractor={(id) => id}
      renderItem={({item: roomId}) => (
        <TouchableOpacity style={styles.item} onPress={() => onCreate()}>
          <Text>{roomId}</Text>
        </TouchableOpacity>
      )}
    />
  );
}

const styles = StyleSheet.create({
  item: {padding: 16, borderBottomWidth: 1},
});
