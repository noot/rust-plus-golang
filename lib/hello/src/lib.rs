#[no_mangle]
pub extern "C" fn init_stuff() {
    env_logger::init();

    log::trace!("init_stuff() trace level message");
    log::debug!("init_stuff() debug level message");
    log::info!("init_stuff() info level message");
    log::warn!("init_stuff() warn level message");
    log::error!("init_stuff() error level message");
}

#[no_mangle]
pub extern "C" fn square_u8_array(
    src: *const libc::c_char,
    arr_len: libc::c_char,
    dst: *mut libc::c_char,
) {
    unsafe {
        let src_vec = std::slice::from_raw_parts::<i8>(src, arr_len as usize);
        let dst_vec: &mut [i8] = std::slice::from_raw_parts_mut::<i8>(dst, arr_len as usize);
        for (i, x) in src_vec.iter().enumerate() {
            dst_vec[i] = *x;
        }
    }
}

// This is present so it's easy to test that the code works natively in Rust via `cargo test
#[cfg(test)]
pub mod test {
    use super::*;

    // This is meant to do the same stuff as the main function in the .go files.
    #[test]
    fn simulated_main_function() {
        init_stuff();
        let src: [i8; 3] = [1, 2, 3];
        let mut dst = [0i8; 3];
        square_u8_array(src.as_ptr(), 3, dst.as_mut_ptr());
    }
}
