import {
    StyleSheet,
    Text,
    TextInput,
    TouchableOpacity,
    View
} from 'react-native';

import Spacer from "../../components/Spacer"
import ThemedText from "../../components/ThemedText"
import ThemedView from "../../components/ThemedView"
import { useRouter} from 'expo-router'
import React, {useCallback, useContext, useMemo, useRef, useState} from "react";
import {useUser} from "../../hooks/useUser";
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import BottomSheet, { BottomSheetView } from '@gorhom/bottom-sheet';
import {UserContext} from "../../contexts/UserContext";

const Feedback = () => {
    const service = useContext(UserContext);
    const { sendFeedBackComments } = useUser()
    const [error, setError] = useState()
    const [inputValue, setInputValue] = useState('');
    const router = useRouter()

    const handleSubmit = async () => {
        setError(null)
        try {
            sendFeedBackComments(service.userId, inputValue)
            router.replace("/reviewing")
        } catch (error) {
            setError(error.message)
        }
    }

    const [text, setText] = useState('');
    // The logic: Button is disabled if text is empty or just whitespace
    const isInvalid = text.trim().length === 0;
    const handleInputChange = (text) => {
        setInputValue(text);
        setText(text)
    };


    const bottomSheetRef = useRef(null);
    const handleSheetChanges = useCallback((index) => {
    }, [])
    const snapPoints = useMemo(() => ['70%', '80%','90%','100%'], []);
    return (
        <GestureHandlerRootView style={styles.container}>
            <BottomSheet
                ref={bottomSheetRef}
                onChange={handleSheetChanges}
                snapPoints={snapPoints}
            >
                <BottomSheetView>
                    <ThemedView style={styles.contentContainer}>
                        <ThemedText style={{fontSize: 25, fontWeight: 'bold'}}>Help us improve Rizon</ThemedText>
                        <View style={{ height: 20  }} />
                        <ThemedText style={styles.feedbackText}>Tell us what didn't feel right, we read every message</ThemedText>
                        <View style={{ height: 20  }} />
                        <TextInput
                            style={styles.input}
                            placeholder="Type your feedback here"
                            placeholderTextColor="gray"
                            value={inputValue}
                            onChangeText={handleInputChange}
                        />
                        <View style={{ height: 20  }} />
                        <TouchableOpacity
                            style={[styles.button, isInvalid && styles.buttonDisabled]}
                            onPress={handleSubmit}
                            disabled={isInvalid}
                        >
                            <Text style={styles.buttonText}>Send feedback</Text>
                        </TouchableOpacity>
                        <Spacer />
                        {error && <Text style={styles.error}>{error}</Text>}
                        <Spacer />
                        <Spacer />
                    </ThemedView>
                </BottomSheetView>
            </BottomSheet>
        </GestureHandlerRootView>

    )
}

export default Feedback

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: "#625f72",
    },
    contentContainer: {
        flex: 1,
        padding: 6,
        alignItems: 'center',
    },

    buttonContainer: {
        flexDirection: 'row',
        flexWrap: 'wrap',
        marginHorizontal: 20,
        marginTop: 5,
    },
    buttonSpacing: {
        marginHorizontal: 5, // Adds 5 units of space to both sides of each button
    },
    feedbackText:{
        fontSize: 14,
        alignSelf: "center",
        fontWeight: 'bold',

    },
    heading: {
        fontWeight: "bold",
        fontSize: 18,
        textAlign: "center",
    },
    input: {
        borderWidth: 1,
        borderColor: '#ccc',
        padding: 18,
        borderRadius: 15,
        marginBottom: 20,
        fontSize: 16,
        width: '80%',
        color: '#0f0f0f',
    },
    button: {
        backgroundColor: '#0f0f0f', // Active color
        padding: 12,
        borderRadius: 20,
        alignItems: 'center',
        marginHorizontal: 3,
        width: '80%'
    },
    buttonDisabled: {
        backgroundColor: '#A9A9A9', // Inactive color (Grey)
    },
    buttonText: {
        color: '#FFFFFF',
        fontWeight: 'bold',
        fontSize: 16,
    },
})
