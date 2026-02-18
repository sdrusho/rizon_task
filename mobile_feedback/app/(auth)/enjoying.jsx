import React, { useCallback, useMemo, useRef,useState,useContext } from 'react';
import { View, Text, StyleSheet,Linking } from 'react-native';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import BottomSheet, { BottomSheetView } from '@gorhom/bottom-sheet';
import ThemedLogo from "../../components/ThemedLogo";
import Spacer from "../../components/Spacer";
import ThemedButton from "../../components/ThemedButton";
import ThemedText from "../../components/ThemedText"
import {router, useRouter} from 'expo-router'
import {useUser} from "../../hooks/useUser";
import {UserContext} from "../../contexts/UserContext";




const Enjoying = () => {
    const service = useContext(UserContext);
    const { sendIsEnjoying} = useUser()
    const router = useRouter()

    const [error, setError] = useState()
    const handleNotyet = async () => {
        setError(null)
        try {
            sendIsEnjoying(service.userId, false)
            router.replace("/feedback")
        } catch (error) {
            setError(error.message)
        }
    }

    const handleLovingIt = async () => {
        setError(null)
        try {
            console.log('current user is enjoying: ', service.userId)
            sendIsEnjoying(service.userId, true)
            router.replace("/feedback")
        } catch (error) {
            setError(error.message)
        }
    }

    const bottomSheetRef = useRef(null);
    const handleSheetChanges = useCallback((index) => {
    }, [])
    const snapPoints = useMemo(() => ['70%', '80%', '90%', '100%'], []);

    return (


        <GestureHandlerRootView style={styles.container}>
            <BottomSheet
                ref={bottomSheetRef}
                onChange={handleSheetChanges}
                snapPoints={snapPoints}
            >
                <BottomSheetView style={styles.contentContainer}>
                    <ThemedLogo/>
                    <Spacer/>

                    <ThemedText style={{fontSize: 25, fontWeight: 'bold'}}>Enjoying Rizon so far?</ThemedText>
                    <View style={{height: 20}}/>
                    <ThemedText style={styles.feedbackText}>Your feedback help us build a better money
                        experience</ThemedText>
                    <View style={{height: 20}}/>
                    <View style={styles.buttonContainer}>

                        <ThemedButton style={{marginHorizontal: 3, width: '40%',}} onPress={handleNotyet}>
                            <Text style={{color: '#f2f2f2', alignSelf: "center"}}>Not yet</Text>
                        </ThemedButton>

                        <ThemedButton style={{marginHorizontal: 3, width: '40%',}} onPress={handleLovingIt}>
                            <Text style={{color: '#f2f2f2', alignSelf: "center"}}>Yes, Loving it</Text>
                        </ThemedButton>


                    </View>
                    <Spacer/>
                    <Spacer/>

                </BottomSheetView>
            </BottomSheet>
        </GestureHandlerRootView>
    );
};
export default Enjoying;

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#625f72',
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
        fontSize: 12,
        alignSelf: "center",
        fontWeight: 'bold',
        color: '#c7c7c7',
    },
});
