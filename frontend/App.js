import React, {useState, useEffect} from 'react';
import {View, Text, TextInput, Button, FlatList, StyleSheet} from 'react-native';
import WebSocketService from './WebSocketService';
import RoomListScreen from './RoomListScreen';
import RoomDetailScreen from './RoomDetailScreen';

export default function App() {
  const [rooms, setRooms] = useState([]);
  const [selectedRoomId, setSelectedRoomId] = useState(null);
  const ws = new WebSocketService();

  useEffect(() => {
    // fetch existing rooms (optional)
    fetch('http://localhost:8080/room')
      .then(res => res.json())
      .then(data => setRooms(prev => [...prev, data.room_id]));
  }, []);

  const handleCreateRoom = async () => {
    const res = await fetch('http://localhost:8080/room', {method: 'POST'});
    if (res.ok) {
      const {room_id} = await res.json();
      setRooms(prev => [...prev, room_id]);
      setSelectedRoomId(room_id);
    }
  };

  return (
    <View style={styles.container}>
      {!selectedRoomId ? (
        <RoomListScreen rooms={rooms} onCreate={handleCreateRoom} />
      ) : (
        <RoomDetailScreen roomId={selectedRoomId} ws={ws} onClose={() => setSelectedRoomId(null)} />
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {flex: 1, backgroundColor: '#fff'},
});
