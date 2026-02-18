import React from 'react';
import {View, Text, StyleSheet} from "react-native";
import ThemedView from '../../components/ThemedView'
import ThemedText  from '../../components/ThemedText';
import { Link } from 'expo-router';

const Success = () => {
    return (
        <View style={styles.container}>
            <View>
                <Text style={{flexShrink: 1}}>
                    Send a mail to your email account {'\n'}
                    Please click the link button for signup complete.
                </Text>
            </View>
        </View>


    );
};
export default Success;

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
        flexDirection: 'row',
        padding: 10
    },

});
