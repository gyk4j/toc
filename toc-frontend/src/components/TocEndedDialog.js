import {
    MDBModal,
    MDBModalDialog,
    MDBModalContent,
    MDBModalHeader,
    MDBModalTitle,
    MDBModalBody,
    MDBModalFooter,
    MDBBtn
} from 'mdb-react-ui-kit';

export default function TocEndedDialog({ 
    showTocEndedDialog,
    toggleShowTocEnded,
    statusbarEndTime,
    actionText
}) {
    return (
        <MDBModal staticBackdrop show={showTocEndedDialog} tabIndex='-1'>
            <MDBModalDialog>
                <MDBModalContent>
                    <MDBModalHeader>
                        <MDBModalTitle>TOC Ended</MDBModalTitle>
                        <MDBBtn className='btn-close' color='none' onClick={toggleShowTocEnded}></MDBBtn>
                    </MDBModalHeader>
                    <MDBModalBody>TOC ended at 
                        <strong className="text-danger"> { statusbarEndTime != null ? statusbarEndTime.toLocaleString() : "" }</strong>
                    </MDBModalBody>
                    <MDBModalFooter>
                        <MDBBtn onClick={toggleShowTocEnded} color='secondary'>
                            {actionText}
                        </MDBBtn>
                    </MDBModalFooter>
                </MDBModalContent>
            </MDBModalDialog>
        </MDBModal>
    )
}
