import React, {useEffect, useState} from 'react';
import {View, Text, TextInput, Button, FlatList, TouchableOpacity, StyleSheet} from 'react-native';
import WebSocketService from './WebSocketService';
import MessageBubble from './MessageBubble';

export default function RoomDetailScreen({roomId, ws, onClose}) {
  const [messages, setMessages] = useState([]);
  const [inputText, setInputText] = useState('');

  useEffect(() => {
    ws.connect(roomId);
    return () => ws.disconnect();
  }, [roomId]);

  const sendMessage = async () => {
    if (!inputText.trim()) return;
    ws.send('message', inputText);
    setInputText('');
  };

  return (
    <View style={styles.container}>
      <FlatList data={messages} renderItem={({item}) => <MessageBubble item={item} />} />
      <TextInput style={styles.input} value={inputText} onChangeText={setInputText} />
      <Button title="Send" onPress={sendMessage} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {flex: 1, padding: 8},
  input: {borderWidth: 1, padding: 8, marginBottom: 8},
});
