import { Pressable, StyleSheet } from 'react-native'
import { Colors } from '../constants/Colors'

function ThemedButton({ style, ...props }) {

    return (
        <Pressable
            style={({ pressed }) => [styles.btn, pressed && styles.pressed, style]}
            {...props}
        />
    )
}

const styles = StyleSheet.create({
    btn: {
        backgroundColor: Colors.button,
        padding: 12,
        borderRadius: 20,
       marginVertical: 10,
        fontSize: 14,
    },
    pressed: {
        opacity: 0.5,
    },
    
})

export default ThemedButton
